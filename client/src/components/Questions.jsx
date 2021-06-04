import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
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
  Container,
  InputBase,
  TextField,
  ThemeProvider,
} from "@material-ui/core";
import Theme from "../styles/Theme";
import useStyles from "../styles/Style";
import HeaderAppBar from "./Header";
import SearchIcon from "@material-ui/icons/Search";

const Questions = () => {
  const classes = useStyles();
  const questionsData = useSelector((state) => state.question);
  const dispatch = useDispatch();
  const [question, setQuestion] = useState([]);
  const [allQuestions, setAllQuestions] = useState([]);
  const [searchQuestions, setSearchQuestions] = useState("");
  const [hidden, setHidden] = useState(false);

  useEffect(() => {
    dispatch(questionAction.fetchAllQuestions());
  }, []);

  useEffect(() => {
    setQuestion(questionsData.data);
    setAllQuestions(questionsData.data);
  }, [question, allQuestions]);

  const filterQuestion = questionsData.data.filter((item) => {
    return searchQuestions !== "" ? item.title.includes(searchQuestions) : "";
  });

  const search = (e) => {
    setSearchQuestions(e.target.value);
    setHidden(true);
  };

  return (
    <>
      <ThemeProvider theme={Theme}>
        <Container
          maxWidth="md"
          component="main"
          className={classes.heroContent}
        >
          <Grid item xs={12} md={12}>
            <Box display="flex" justifyContent="space-between">
              <div>
                <form>
                  <TextField
                    placeholder="Search"
                    type="text"
                    size="small"
                    variant="outlined"
                    onChange={search}
                  />
                </form>
              </div>
              <Typography component="h2" variant="h5">
                Questions
              </Typography>
              <Button
                href="/questions/ask"
                color="secondary"
                variant="contained"
              >
                Ask Question
              </Button>
            </Box>
            {searchQuestions == ""
              ? questionsData.data.map((data, index) => {
                  return (
                    <Card className={classes.card} key={index}>
                      <div className={classes.cardDetails}>
                        <CardContent className={classes.cardContent}>
                          <Link to={`/questions/${data.id}`}>
                            <Typography component="h2" variant="h6">
                              {data.title}
                            </Typography>
                          </Link>
                          <Typography
                            variant="body2"
                            paragraph
                            className={classes.customBox}
                          >
                            {data.content}
                          </Typography>
                          <Typography variant="caption" color="primary">
                            {data.category_name}
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
                              <Avatar alt={data.user_name} src={data.avatar} />
                              <Typography
                                variant="subtitle2"
                                color="textSecondary"
                              >
                                <Link
                                  href="#"
                                  style={{
                                    marginLeft: "1rem",
                                  }}
                                >
                                  {data.user_name}
                                </Link>
                              </Typography>
                            </Box>
                          </Box>
                        </CardContent>
                      </div>
                    </Card>
                  );
                })
              : null}
            {filterQuestion &&
              filterQuestion.map((data, index) => {
                return (
                  <Card className={classes.card} key={index}>
                    <div className={classes.cardDetails}>
                      <CardContent className={classes.cardContent}>
                        <Link to={`/questions/${data.id}`}>
                          <Typography component="h2" variant="h6">
                            {data.title}
                          </Typography>
                        </Link>
                        <Typography
                          variant="body2"
                          paragraph
                          className={classes.customBox}
                        >
                          {data.content}
                        </Typography>
                        <Typography variant="caption" color="primary">
                          {data.category_name}
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
                            <Avatar alt={data.user_name} src={data.avatar} />
                            <Typography
                              variant="subtitle2"
                              color="textSecondary"
                            >
                              <Link
                                href="#"
                                style={{
                                  marginLeft: "1rem",
                                }}
                              >
                                {data.user_name}
                              </Link>
                            </Typography>
                          </Box>
                        </Box>
                      </CardContent>
                    </div>
                  </Card>
                );
              })}
          </Grid>
        </Container>
      </ThemeProvider>
    </>
  );
};

export default Questions;
