package repositories

import (
	"api/src/models"
	"database/sql"
)

type Publications struct {
	db *sql.DB
}

func PublicationRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

func (repository Publications) Create(publication models.Publication) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into publications (title, content, author_id) values(?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Content, publication.AuthorID)
	if err != nil {
		return 0, err
	}

	lastSavedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastSavedID), nil
}

func (repository Publications) SearchByID(publicationID uint64) (models.Publication, error) {
	line, err := repository.db.Query(
		`select p.*, u.nick 
		from publications p inner join users u on u.id = p.author_id 
		where p.id = ?`,
		publicationID,
	)
	if err != nil {
		return models.Publication{}, err
	}
	defer line.Close()

	var publication models.Publication

	if line.Next() {
		if err = line.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNickName,
		); err != nil {
			return models.Publication{}, nil
		}
	}

	return publication, nil
}

func (repository Publications) SearchPublications(userID uint64) ([]models.Publication, error) {
	lines, err := repository.db.Query(`
		select distinct p.*, u.nick_name 
		from publications p 
		inner join users u on u.id = p.author_id 
		inner join followers f on p.author_id = f.user_id 
		where u.id = ? or f.follow_id = ? 
		order by 1 desc`,
		userID,
		userID,
	)
	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication

		if err = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNickName,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (repository Publications) Update(publicationID uint64, publication models.Publication) error {
	statement, err := repository.db.Prepare("update publications set title = ?, content = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publication.Title, publication.Content, publicationID); err != nil {
		return err
	}
	return nil
}

func (repository Publications) Delete(publicationID uint64) error {
	statement, err := repository.db.Prepare("delete from publications where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}
	return nil
}

func (repository Publications) SearchPublicationsByUserID(userID uint64) ([]models.Publication, error) {
	lines, err := repository.db.Query(`
		select distinct p.*, u.nick_name 
		from publications p 
		inner join users u on u.id = p.author_id 
		where p.author_id = ?`,
		userID,
	)
	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication

		if err = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNickName,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (repository Publications) LikeIt(publicationID uint64) error {
	statement, err := repository.db.Prepare("update publications set likes = likes + 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}

func (repository Publications) DislikeIt(publicationID uint64) error {
	statement, err := repository.db.Prepare(`
		update publications set likes = 
		CASE 
			WHEN likes > 0 THEN likes - 1
			ELSE likes 
		END
		where id = ?`,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}
