//import TextField from "@material-ui/core/TextField";
import React, {useState} from "react";
import "./searchBar.css"
import SearchIcon from '@mui/icons-material/Search';
import CloseIcon from '@mui/icons-material/Close';
/*import Package from "../package-lock.json"
const searchBar = () => {}
const [searchInput, setSearchInput] = useState("")*/

//Implementation here
function Search({placeholder, data}){
    const [filteredData, setFilteredData] = useState([]); //constant states
    const [wordEntered, setWordEntered] = useState("");  //constant states
    const handleFilter = (event) => {
        const searchWord = event.target.value;  //word that user is trying to search for
        setWordEntered(searchWord);
        const newFilter = data.filter((value) => {
            return value.title.toLowerCase().includes(searchWord.toLowerCase());
        });

        if (searchWord === ""){
            setFilteredData([]);
        } else {
            setFilteredData(newFilter);
        }
    };
    const clearInput = () => {
        setFilteredData([]);
        setWordEntered("");
    }
    return (
        <div className="search">
            <div className="searchInputs">
                <input type="text" placeholder={placeholder} value={wordEntered} onChange={handleFilter}/>
                <div className="SearchIcon">
                    {filteredData.length === 0 ? (
                        <SearchIcon />
                    ): (
                        <CloseIcon id="clearBtn" onClick={clearInput} />
                    )}
                </div>
            </div>
            {filteredData.length != 0 && (
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