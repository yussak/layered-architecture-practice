import { useEffect, useState } from 'react'
import './App.css'

function App() {
  const [todos, setTodos] = useState([]);

  useEffect(() => {
    fetch('http://localhost:8080/')
      .then((response) => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then((data) => setTodos(data))
      .catch((error) => console.error('Error fetching data:', error));
  }, []);

  return (
    <>
      <h1>TODOリスト</h1>
      <form action="/add" method="POST">
        <input type="text" name="todo" placeholder="タスクを入力" required />
        <button type="submit">追加</button>
      </form>
      <ul>
        {todos && todos.length > 0 ? (
          todos.map((todo) => (
            <li key={todo.id}>
              {todo.name} <a href={`/delete?id=${todo.id}`}>削除</a>
            </li>
          ))
        ) : (
          <p>Todoがありません</p>
        )}
      </ul>
    </>
  )
}

export default App
