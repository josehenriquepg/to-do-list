# To-Do List

This is a command-line application written in Go (Golang) for managing a to-do list. It allows you to add, list, complete, and remove tasks, with file persistence support to save tasks between sessions.

## :books: Features

- **Add Task**: Allows you to add a new task.
- **List Tasks**: Displays all pending and completed tasks.
- **Complete Task**: Marks a task as completed.
- **Remove Task**: Removes a task from the list.
- **File Persistence**: Saves the task list to a JSON file and automatically loads it on startup.

## :floppy_disk: Data Structure

The data structure for a task is as follows:
```go
type Task struct {
    ID          int
    Description string
    Completed   bool
}
```

## :gear: How to Use

#### Prerequisites
  - Go 1.16+ installed on your system.
#### Installation
  1. Clone the repository:

```bash
git clone https://github.com/josehenriquepg/to-do-list.git
cd to-do-list
```

  2. Build the project:

```bash
go build -o to-do-list
```

#### Running the Application
Run the compiled executable:

```bash
./to-do-list
```

You will see a menu with the following options:

```markdown
Menu:
1. Add Task
2. List Tasks
3. Complete Task
4. Remove Task
5. Exit
Choose an option:
```

#### Usage Example
1. Add a Task:

  - Choose option 1.
  - Enter the task description when prompted.

2. List Tasks:

  - Choose option 2 to view all tasks with their statuses.

3. Complete a Task:

  - Choose option 3.
  - Enter the ID of the task you want to mark as completed.

4. Remove a Task:

  - Choose option 4.
  - Enter the ID of the task you want to remove.

#### Code Structure
  - **main.go**: Contains the main logic of the application, including functions to add, list, complete, and remove tasks, as well as functions to save and load tasks from a file.

## :handshake: Collaborators 

<table>
  <tr>
    <td align="center">
      <a href="http://github.com/josehenriquepg">
        <img src="https://avatars.githubusercontent.com/josehenriquepg" width="100px;" alt="Foto de José Henrique no GitHub"/><br>
        <sub>
          <b>José Henrique</b>
        </sub>
      </a>
    </td>
  </tr>
</table>