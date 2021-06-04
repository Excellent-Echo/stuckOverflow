import {
  Box,
  Button,
  Card,
  CardContent,
  CircularProgress,
  Container,
  CssBaseline,
  Grid,
  Link,
  TextField,
  Typography,
} from "@material-ui/core";
import { ThemeProvider } from "@material-ui/styles";
import CloudUploadIcon from "@material-ui/icons/CloudUpload";
import React, { useEffect, useState } from "react";
import PrimarySearchAppBar from "../components/Header";
import useStyles from "../styles/Style";
import Theme from "../styles/Theme";
import { useDispatch, useSelector } from "react-redux";
import { useHistory, useParams } from "react-router";
import postQuestionAction from "../redux/question/post/postQuestionAction";
import Copyright from "../components/Copyright";
import Alert from "@material-ui/lab/Alert";
import Header from "../components/Header";
import detailQuestionAction from "../redux/question/detail/detailQuestionAction";

const QuestionDetail = () => {
  const classes = useStyles();
  const detailQuestionsData = useSelector((state) => state.detailQuestion);
  const dispatch = useDispatch();
  const history = useHistory();
  const { id } = useParams();

  useEffect(() => {
    if (id && id !== "") {
      dispatch(detailQuestionAction.fetchDetailQuestion(id));
    }

    return () => {
      dispatch(detailQuestionAction.removeDetailQuestion());
    };
  }, [id]);

  return (
    <div>
      <ThemeProvider theme={Theme}>
        {Object.keys(detailQuestionsData).length === 0 ? (
          <Box
            display="flex"
            justifyContent="center"
            alignItems="center"
            minHeight="100vh"
          >
            <CircularProgress />
          </Box>
        ) : (
          <div>
            <CssBaseline />
            <Header />
            <Container
              maxWidth="md"
              component="main"
              className={classes.heroContent}
            >
              <Grid item xs={12} md={12}>
                <Box display="flex" justifyContent="flex-end">
                  <Button
                    href="/questions/ask"
                    color="secondary"
                    variant="contained"
                  >
                    Ask Question
                  </Button>
                </Box>
              </Grid>
              <h1>{detailQuestionsData.data.data.title}</h1>
              {/* <pre>{JSON.stringify(detailQuestionsData.data.data)}</pre> */}
              <p>{detailQuestionsData.data.data.content}</p>
            </Container>
            <Container
              maxWidth="md"
              component="footer"
              className={classes.footer}
            >
              <Box mt={5}>
                <Copyright />
              </Box>
            </Container>
          </div>
        )}
      </ThemeProvider>
    </div>
  );
};

export default QuestionDetail;
