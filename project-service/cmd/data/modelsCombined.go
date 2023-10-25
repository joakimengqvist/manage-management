package data

import (
	"context"
)

func AppendProjectsToSubProject(projectIds []string, subProjectId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `
		with new_projects as (
			select distinct unnest(projects || $1) as value
			from sub_projects
			where id = $2
		)
		update sub_projects
		set projects = array(select value from new_projects)
		where id = $2
	`

	_, err := db.ExecContext(ctx, stmt, convertToPostgresArray(projectIds), subProjectId)
	if err != nil {
		return err
	}

	for _, projectId := range projectIds {
		stmt := `update projects set sub_projects = array(SELECT DISTINCT unnest(array_append(sub_projects, $1))) where id = $2`

		_, err := db.ExecContext(ctx, stmt, subProjectId, projectId)
		if err != nil {
			return err
		}
	}

	return nil
}

func AppendSubProjectsToProject(projectId string, subProjectIds []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `
		with new_sub_projects as (
			select distinct unnest(sub_projects || $1) as value
			from projects
			where id = $2
		)
		update projects
		set sub_projects = array(select value from new_sub_projects)
		where id = $2
	`

	_, err := db.ExecContext(ctx, stmt, convertToPostgresArray(subProjectIds), projectId)
	if err != nil {
		return err
	}

	for _, subProjectId := range subProjectIds {
		stmt := `update sub_projects set projects = array(SELECT DISTINCT unnest(array_append(projects, $1))) where id = $2`

		_, err := db.ExecContext(ctx, stmt, projectId, subProjectId)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteProjectsFromSubProject(projectIds []string, subProjectId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	for _, projectId := range projectIds {

		stmtUpdateSubProjects := `update sub_projects set projects = array_remove(projects, $1) where id = $2`

		_, err := db.ExecContext(ctx, stmtUpdateSubProjects, projectId, subProjectId)
		if err != nil {
			return err
		}

		stmtUpdateProjects := `update projects set sub_projects = array_remove(sub_projects, $1) where id = $2`

		_, err = db.ExecContext(ctx, stmtUpdateProjects, subProjectId, projectId)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteSubProjectsFromProject(projectId string, subProjectIds []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	for _, subProjectId := range subProjectIds {

		stmtUpdateSubProjects := `update projects set sub_projects = array_remove(sub_projects, $1) where id = $2`

		_, err := db.ExecContext(ctx, stmtUpdateSubProjects, subProjectId, projectId)
		if err != nil {
			return err
		}

		stmtUpdateProjects := `update sub_projects set projects = array_remove(projects, $1) where id = $2`

		_, err = db.ExecContext(ctx, stmtUpdateProjects, projectId, subProjectId)
		if err != nil {
			return err
		}
	}
	return nil
}
