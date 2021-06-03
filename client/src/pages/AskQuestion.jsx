import {
  Box,
  Button,
  Card,
  CardContent,
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
import { useHistory } from "react-router";
import postQuestionAction from "../redux/question/post/postQuestionAction";
import Copyright from "../components/Copyright";

const AskQuestion = () => {
  const classes = useStyles();
  const postQuestionData = useSelector((state) => state.postQuestion);
  const dispatch = useDispatch();
  const history = useHistory();
  const isAuthUser = !!localStorage.getItem("accessToken");
  const [validated, setValidated] = useState(false);

  useEffect(() => {
    if (!isAuthUser) {
      history.push("/login");
    }
    dispatch(postQuestionAction.resetForm());
  }, []);

  const handleSubmitQuestion = (e) => {
    e.preventDefault();
    const form = e.currentTarget;
    if (form.checkValidity() === false) {
      e.preventDefault();
      e.stopPropagation();
    } else {
      dispatch(
        postQuestionAction.askQuestion(
          postQuestionData.title,
          postQuestionData.content,
          postQuestionData.categoryID
        )
      );
    }
    setValidated(true);
    console.log(postQuestionData);
  };

  return (
    <>
      <ThemeProvider theme={Theme}>
        <PrimarySearchAppBar />
        <CssBaseline />
        <Container
          maxWidth="md"
          component="main"
          className={classes.heroContent}
        >
          <Typography component="h2" variant="h5">
            Ask a question
          </Typography>
          <Card className={classes.card}>
            <div className={classes.cardDetails}>
              <CardContent className={classes.cardContent}>
                <form
                  className={classes.form}
                  noValidate
                  onSubmit={handleSubmitQuestion}
                >
                  <Grid container spacing={2}>
                    <Grid item xs={12}>
                      <Typography variant="h6">Title</Typography>
                      <Typography variant="caption">
                        Be specific and imagine you’re asking a question to
                        another person
                      </Typography>

                      <TextField
                        placeholder="e.g. Is there an R function for finding the index of en element in a vector?"
                        variant="outlined"
                        fullWidth
                        required
                        value={postQuestionData.title}
                        onChange={(e) =>
                          dispatch(postQuestionAction.setTitle(e.target.value))
                        }
                      />
                    </Grid>
                    <Grid item xs={12}>
                      <Typography variant="h6">Body</Typography>
                      <Typography variant="caption">
                        Include all the information someone would need to answer
                        your question
                      </Typography>
                      <TextField
                        multiline={true}
                        rows={7}
                        size="medium"
                        variant="outlined"
                        fullWidth
                        required
                        value={postQuestionData.content}
                        onChange={(e) =>
                          dispatch(
                            postQuestionAction.setContent(e.target.value)
                          )
                        }
                      />
                    </Grid>
                    <Grid item xs={12}>
                      <Button
                        type="submit"
                        variant="contained"
                        color="default"
                        className={classes.button}
                        startIcon={<CloudUploadIcon />}
                      >
                        Add an attachment
                      </Button>
                    </Grid>
                    <Grid item xs={12}>
                      <Typography variant="h6">Category</Typography>
                      <Typography variant="caption">
                        Add category to describe what your question is about
                      </Typography>
                      <TextField
                        size="small"
                        variant="outlined"
                        fullWidth
                        required
                        value={postQuestionData.categoryID}
                        onChange={(e) =>
                          dispatch(
                            postQuestionAction.setCategory(e.target.value)
                          )
                        }
                      />
                    </Grid>
                  </Grid>
                  <Button
                    type="submit"
                    fullWidth
                    variant="contained"
                    color="secondary"
                    disabled={postQuestionData.isLoading ? true : false}
                  >
                    {postQuestionData.isLoading ? "Loading..." : "Submit"}
                  </Button>
                </form>
              </CardContent>
            </div>
          </Card>
        </Container>
        {/* Footer */}
        <Container maxWidth="md" component="footer" className={classes.footer}>
          <Box mt={5}>
            <Copyright />
          </Box>
        </Container>
      </ThemeProvider>
    </>
  );
};

export default AskQuestion;
