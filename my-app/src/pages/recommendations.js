import data from "./MusicSheet.json";
import React, {useState} from 'react'
import "./recommendations.css"

const Recommendations = () =>{
    const [userInput, setUserInput] = useState("")
    const [recommendations, userRecommendations] = useState([])
    //goes through json database to filter songs based on user input - event handler
    const handleRecommend = () => {
        const filteredSongs = data.filter((song) =>
          song.genre.toLowerCase().includes(userInput.toLowerCase())
        );
        userRecommendations(filteredSongs);
      };

    return (
        <div>
        <h1> Song Recommendations</h1>
        <input 
        type="text"
        value = {userInput}
        //onChange - event handler
        onChange={(e) => setUserInput(e.target.value)}
        placeholder = "Tell me a genre you want"
        />
        <button onClick={handleRecommend}>Recommend songs</button>
      {recommendations.length > 0 && (
        <div>
          <h2>Recommended Songs within Playlist:</h2>
          <ul>
            {recommendations.map((song, index) => (
              <li key={index}>
                <a href={song.link} target="_blank" rel="noopener noreferrer">
                  {song.song} by {song.artist} ({song.year}) - {song.genre}
                </a>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
    );
};

export default Recommendations;
