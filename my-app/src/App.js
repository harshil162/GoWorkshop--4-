import music from './music.png';
import './App.css';
import "./searchBar.js"
function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={music} className="App-logo" alt="music" />
        <p>
          Welcome to my Music Downloader App.
        </p>
        <a
          className="App-link"
          //href="https://reactjs.org"
          href="localhost:3000"
          target="_blank"
          rel="noopener noreferrer"
        >
          Use This App
        </a>
      </header>
    </div>
  );
}

export default App;
