package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func TestCreateTask(t *testing.T) {
	t.Cleanup(func() { cleanupDatabase(t) })

	task, err := testQueries.CreateTask(ctx, "do test")
	require.NoError(t, err)

	require.Equal(t, "do test", task.Description)
	require.Equal(t, StatusTodo, task.Status)
	require.NotNil(t, task.ID)
	require.NotNil(t, task.CreatedAt)
	require.Zero(t, task.UpdatedAt)
	require.Zero(t, task.DeletedAt)
}

func TestDeleteTask(t *testing.T) {
	t.Cleanup(func() { cleanupDatabase(t) })

	task, err := testQueries.CreateTask(ctx, "do test")
	require.NoError(t, err)

	err = testQueries.DeleteTask(ctx, task.ID)
	require.NoError(t, err)

	tasks, err := testQueries.ListTasks(ctx, ListTasksParams{
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Len(t, tasks, 0)
}

func TestGetTaskForUpdate(t *testing.T) {
	t.Cleanup(func() { cleanupDatabase(t) })

	expectedTask, err := testQueries.CreateTask(ctx, "do test")
	require.NoError(t, err)

	actualTask, err := testQueries.GetTaskForUpdate(ctx, expectedTask.ID)
	require.NoError(t, err)

	require.Equal(t, expectedTask, actualTask)
}

func TestListTasks(t *testing.T) {
	t.Cleanup(func() { cleanupDatabase(t) })

	taskDescriptions := []string{"task 1", "task 2", "task 3"}
	for _, desc := range taskDescriptions {
		_, err := testQueries.CreateTask(ctx, desc)
		require.NoError(t, err)
	}

	tasks, err := testQueries.ListTasks(ctx, ListTasksParams{
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)

	require.Len(t, tasks, len(taskDescriptions))

	for i, task := range tasks {
		require.Equal(t, taskDescriptions[i], task.Description)
		require.Equal(t, StatusTodo, task.Status)
		require.NotNil(t, task.ID)
		require.NotNil(t, task.CreatedAt)
		require.Zero(t, task.UpdatedAt)
		require.Zero(t, task.DeletedAt)
	}
}

func TestListTasksByStatus(t *testing.T) {
	t.Cleanup(func() { cleanupDatabase(t) })

	taskDescriptions := []string{"task 1", "task 2", "task 3"}
	for _, desc := range taskDescriptions {
		_, err := testQueries.CreateTask(ctx, desc)
		require.NoError(t, err)
	}

	tasks, err := testQueries.ListTasksByStatus(ctx, ListTasksByStatusParams{
		Limit:  10,
		Offset: 0,
		Status: StatusTodo,
	})
	require.NoError(t, err)

	require.Len(t, tasks, len(taskDescriptions))

	for i, task := range tasks {
		require.Equal(t, taskDescriptions[i], task.Description)
		require.Equal(t, StatusTodo, task.Status)
		require.NotNil(t, task.ID)
		require.NotNil(t, task.CreatedAt)
		require.Zero(t, task.UpdatedAt)
		require.Zero(t, task.DeletedAt)
	}
}

func TestUpdateTaskDescription(t *testing.T) {
	t.Cleanup(func() { cleanupDatabase(t) })

	task, err := testQueries.CreateTask(ctx, "initial description")
	require.NoError(t, err)

	newDescription := "updated description"
	updatedTask, err := testQueries.UpdateTaskDescription(ctx, UpdateTaskDescriptionParams{
		ID:          task.ID,
		Description: newDescription,
	})
	require.NoError(t, err)

	require.Equal(t, newDescription, updatedTask.Description)
	require.Equal(t, task.ID, updatedTask.ID)
	require.Equal(t, StatusTodo, updatedTask.Status)
	require.NotNil(t, updatedTask.CreatedAt)
	require.NotNil(t, updatedTask.UpdatedAt)
	require.Zero(t, updatedTask.DeletedAt)
}

func TestUpdateTaskStatus(t *testing.T) {
	t.Cleanup(func() { cleanupDatabase(t) })

	task, err := testQueries.CreateTask(ctx, "task to update status")
	require.NoError(t, err)

	newStatus := StatusInProgress
	updatedTask, err := testQueries.UpdateTaskStatus(ctx, UpdateTaskStatusParams{
		ID:     task.ID,
		Status: newStatus,
	})
	require.NoError(t, err)

	require.Equal(t, newStatus, updatedTask.Status)
	require.Equal(t, task.ID, updatedTask.ID)
	require.NotNil(t, updatedTask.CreatedAt)
	require.NotNil(t, updatedTask.UpdatedAt)
	require.Zero(t, updatedTask.DeletedAt)
}

func cleanupDatabase(t testing.TB) {
	t.Helper()
	if _, err := testDB.Exec("TRUNCATE TABLE task CASCADE"); err != nil {
		t.Fatalf("failed to clean up after test: %v", err)
	}
}
