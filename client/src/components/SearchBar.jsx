import { InputBase } from "@material-ui/core";
import React, { useState } from "react";
import useStyles from "../styles/Style";
import SearchIcon from "@material-ui/icons/Search";
import { useDispatch } from "react-redux";
import searchQuestionAction from "../redux/question/search/searchQuestionAction";
import { useHistory } from "react-router";

const SearchBar = ({ onChange }) => {
  const [querySearch, setQuerySearch] = useState("");
  const classes = useStyles();
  const dispatch = useDispatch();
  const history = useHistory();

  const handleSearch = (e) => {
    setQuerySearch(e.target.value);
  };

  const handleSubmitSearch = () => {
    dispatch(searchQuestionAction.fetchSearchQuestions(querySearch));
    history.push({
      pathname: "/questions/result/" + querySearch,
    });
  };

  return (
    <>
      <div className={classes.search}>
        <div className={classes.searchIcon}>
          <SearchIcon />
        </div>
        <form onSubmit={handleSubmitSearch}>
          <InputBase
            placeholder="Search…"
            classes={{
              root: classes.inputRoot,
              input: classes.inputInput,
            }}
            inputProps={{ "aria-label": "search" }}
            onChange={handleSearch}
          />
        </form>
      </div>
    </>
  );
};

export default SearchBar;
