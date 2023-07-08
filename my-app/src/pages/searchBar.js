//import TextField from "@material-ui/core/TextField";
import React, {useState} from "react";
import "./searchBar.css"
import SearchIcon from '@mui/icons-material/Search';
import CloseIcon from '@mui/icons-material/Close';
import "./MusicSheet.json";
/*import Package from "../package-lock.json"
const searchBar = () => {}
const [searchInput, setSearchInput] = useState("")*/

//Implementation here
function Search({placeholder, data}){
    const [filteredData, setFilteredData] = useState([]); //constant states
    const [nameEntered, setNameEntered] = useState("");  //constant states
    const handleFilter = (event) => {
        const searchSong = event.target.value;  //word that user is trying to search for
        setNameEntered(searchSong);
        const newFilter = data.filter((value) => {
            return value.song.toLowerCase().includes(searchSong.toLowerCase());
        });

        if (searchSong === ""){
            setFilteredData([]);
        } else {
            setFilteredData(newFilter);
        }
    };
    const clearInput = () => {
        setFilteredData([]);
        setNameEntered("");
    }
    return (
        <div className="search">
            <div className="searchInputs">
                <input type="text" placeholder={placeholder} value={nameEntered} onChange={handleFilter}/>
                <div className="SearchIcon">
                    {filteredData.length === 0 ? (
                        <SearchIcon />
                    ): (
                        <CloseIcon id="clearBtn" onClick={clearInput} />
                    )}
                </div>
            </div>
            {filteredData.length !== 0 && (
            <div className="dataResult">
                {filteredData.slice(0,15).map((value, key)=> {
                    return (
                    <a className="dataItem" href={value.link} target="_blank">
                        <p>{value.song}</p>
                    </a>
                    );
                })}
            </div>
            )}
        </div>
    )
}
export default Search;