import React, { useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Grid from "@material-ui/core/Grid";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardMedia from "@material-ui/core/CardMedia";
import Hidden from "@material-ui/core/Hidden";
import { useDispatch, useSelector } from "react-redux";
import questionAction from "../redux/question/questionAction";
import LoadingSpinner from "./Spinner";
import {
  Avatar,
  Box,
  Button,
  CircularProgress,
  Link,
  ThemeProvider,
} from "@material-ui/core";
import Theme from "../styles/Theme";
import useStyles from "../styles/Style";

const Questions = () => {
  const classes = useStyles();
  const questionsData = useSelector((state) => state.question);
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(questionAction.fetchAllQuestions());
  }, []);

  return (
    <>
      <ThemeProvider theme={Theme}>
        <Grid item xs={12} md={12}>
          <Box display="flex" justifyContent="space-between">
            <Typography component="h2" variant="h5">
              Questions
            </Typography>
            <Button href="/questions/ask" color="secondary" variant="contained">
              Ask Question
            </Button>
          </Box>
          {questionsData.data &&
            questionsData.data.map((data, index) => {
              return (
                <Card className={classes.card} key={index}>
                  <div className={classes.cardDetails}>
                    <CardContent className={classes.cardContent}>
                      <Typography component="h2" variant="h6">
                        {data.title}
                      </Typography>
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
                          <Typography variant="subtitle2" color="textSecondary">
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
      </ThemeProvider>
    </>
  );
};

export default Questions;
