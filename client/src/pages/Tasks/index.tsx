import React from 'react'
import { Button } from 'antd'
import { useHistory, useParams } from 'react-router-dom'
import 'antd/dist/antd.css'
import './index.scss'

function App() {
  const params = useParams()
  const {id} = useParams()
  const history = useHistory()

  return (
    <div className="App">
      <header className="App-header">
        <p>这里是App</p>
        <br />
        <p>Task ID: {JSON.stringify(params)}</p>
        <br />
        <p>{id}</p>
        <Button onClick={() => {
          history.push('/user/home')
        }}>返回工作台</Button>
      </header>
    </div>
  );
}

export default App;
