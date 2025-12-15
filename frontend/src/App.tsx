import { useState } from 'react'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="App">
      <header className="App-header">
        <h1>testlikexin-12-15-demo</h1>
        <p>Frontend for testlikexin-12-15-demo</p>
        
        <div className="card">
          <button onClick={() => setCount((count) => count + 1)}>
            计数: {count}
          </button>
        </div>

        <div className="info">
          <p>
            编辑 <code>src/App.tsx</code> 并保存以查看热重载效果
          </p>
        </div>

        <div className="links">
          <a href="https://react.dev" target="_blank" rel="noopener noreferrer">
            React 文档
          </a>
          <a href="https://vitejs.dev" target="_blank" rel="noopener noreferrer">
            Vite 文档
          </a>
          <a href="https://www.typescriptlang.org/" target="_blank" rel="noopener noreferrer">
            TypeScript 文档
          </a>
        </div>
      </header>
    </div>
  )
}

export default App

