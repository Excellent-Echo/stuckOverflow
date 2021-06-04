import React, { useEffect, useState } from "react";
import { Link, useHistory } from "react-router-dom";
import { makeStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Grid from "@material-ui/core/Grid";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardMedia from "@material-ui/core/CardMedia";
import Hidden from "@material-ui/core/Hidden";
import { useDispatch, useSelector } from "react-redux";
import questionAction from "../redux/question/questionAction";
import {
  Avatar,
  Box,
  Button,
  CircularProgress,
  ThemeProvider,
} from "@material-ui/core";
import Theme from "../styles/Theme";
import useStyles from "../styles/Style";
import HeaderAppBar from "../components/Header";
import searchQuestionAction from "../redux/question/search/searchQuestionAction";

const Search = () => {
  const classes = useStyles();
  const questionsData = useSelector((state) => state.question);
  const dispatch = useDispatch();
  const [allQuestions, setAllQuestions] = useState(questionsData);
  const [searchQuestions, setSearchQuestions] = useState("");
  // const [query, setQuery] = useState("");
  const history = useHistory();
  const searchData = useSelector((state) => state.searchQuestion);

  let search = window.location.search;
  let params = new URLSearchParams(search);
  let foo = params.get("search");

  useEffect(() => {
    dispatch(searchQuestionAction.fetchSearchQuestions(foo));
    console.log(window.location.search);
    console.log(searchData);
    console.log(foo);
  }, []);

  return (
    <>
      <ThemeProvider theme={Theme}>
        <HeaderAppBar />
        <h1></h1>
        <Grid item xs={12} md={12}>
          <Box display="flex" justifyContent="space-between">
            <Typography component="h2" variant="h5">
              Questions
            </Typography>
            <Button href="/questions/ask" color="secondary" variant="contained">
              Ask Question
            </Button>
          </Box>
          {/* {searchData.data && */}
          {/* searchData.data.map((data, index) => { */}
          {/* return ( */}
          <Card className={classes.card}>
            <div className={classes.cardDetails}>
              <CardContent className={classes.cardContent}>
                {/* <Link to={`/questions/${data.id}`}> */}
                <Typography component="h2" variant="h6">
                  Title
                </Typography>
                {/* </Link> */}
                <Typography
                  variant="body2"
                  paragraph
                  className={classes.customBox}
                >
                  Content
                </Typography>
                <Typography variant="caption" color="primary">
                  category
                </Typography>
                <Box
                  display="flex"
                  flexDirection="column"
                  alignItems="flex-end"
                >
                  <Typography variant="caption" color="textSecondary">
                    asked by
                  </Typography>
                  <Box display="flex" color="textSecondary">
                    {/* <Avatar alt={data.user_name} src={data.avatar} /> */}
                    <Typography variant="subtitle2" color="textSecondary">
                      <Link
                        href="#"
                        style={{
                          marginLeft: "1rem",
                        }}
                      >
                        username
                      </Link>
                    </Typography>
                  </Box>
                </Box>
              </CardContent>
            </div>
          </Card>
          {/* ); */}
          {/* })} */}
        </Grid>
      </ThemeProvider>
    </>
  );
};

export default Search;
