import './App.css'
import URLShortener from './components/URLShortener'

function App() {
  return (
    <div className="app-container">
      <header>
        <h1>短链接服务</h1>
        <p>将长链接转换为简短易记的URL</p>
      </header>
      
      <main>
        <URLShortener />
      </main>
      
      <footer>
        <p>© {new Date().getFullYear()} 短链接服务</p>
      </footer>
    </div>
  )
}

export default App

