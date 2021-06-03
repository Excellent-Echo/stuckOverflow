import {
  Box,
  Button,
  Card,
  CardContent,
  CircularProgress,
  Container,
  CssBaseline,
  Grid,
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
import Alert from "@material-ui/lab/Alert";
import categoryAction from "../redux/category/get/categoryAction";
import Autocomplete from "@material-ui/lab/Autocomplete";

const AskQuestion = () => {
  const classes = useStyles();
  const postQuestionData = useSelector((state) => state.postQuestion);
  const categoryList = useSelector((state) => state.categoryList);
  const dispatch = useDispatch();
  const history = useHistory();
  const isAuthUser = !!localStorage.getItem("accessToken");
  const [value, setValue] = useState("");
  const [inputValue, setInputValue] = useState("");

  useEffect(() => {
    dispatch(categoryAction.fetchCategoryList());
    if (!isAuthUser) {
      history.push("/login");
    }
    dispatch(postQuestionAction.resetForm());
  }, []);

  const handleSubmitQuestion = (e) => {
    e.preventDefault();
    const form = e.currentTarget;
    const categoryID = value.id;
    if (form.checkValidity() === false) {
      e.preventDefault();
      e.stopPropagation();
    } else {
      dispatch(
        postQuestionAction.askQuestion(
          postQuestionData.title,
          postQuestionData.content,
          categoryID,
          history
        )
      );
    }
    setValue(categoryList.data[0]);
    setInputValue(value.category_name);
  };

  return (
    <div>
      {Object.keys(categoryList).length === 0 ? (
        <Box
          display="flex"
          justifyContent="center"
          alignItems="center"
          minHeight="100vh"
        >
          <CircularProgress />
        </Box>
      ) : (
        <ThemeProvider theme={Theme}>
          <PrimarySearchAppBar />
          <CssBaseline />
          {postQuestionData.errorMessage && (
            <Alert severity="error">{postQuestionData.errorMessage}</Alert>
          )}
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
                            dispatch(
                              postQuestionAction.setTitle(e.target.value)
                            )
                          }
                        />
                      </Grid>
                      <Grid item xs={12}>
                        <Typography variant="h6">Body</Typography>
                        <Typography variant="caption">
                          Include all the information someone would need to
                          answer your question
                        </Typography>
                        <TextField
                          type="text"
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
                        <Autocomplete
                          id="combo-box-demo"
                          value={value}
                          inputValue={inputValue}
                          onChange={(event, newValue) => {
                            setValue(newValue);
                          }}
                          onInputChange={(event, newInputValue) => {
                            setInputValue(newInputValue);
                          }}
                          options={categoryList.data}
                          getOptionLabel={(option) => option.category_name}
                          getOptionSelected={(option, value) =>
                            option.id === value.id
                          }
                          renderInput={(params) => (
                            <TextField
                              {...params}
                              fullWidth
                              required
                              variant="outlined"
                            />
                          )}
                        />
                      </Grid>
                    </Grid>
                    <Box mt={5} display="flex" justifyContent="center">
                      <Button
                        type="submit"
                        variant="contained"
                        color="secondary"
                        disabled={postQuestionData.isLoading ? true : false}
                      >
                        {postQuestionData.isLoading ? "Loading..." : "Submit"}
                      </Button>
                    </Box>
                  </form>
                </CardContent>
              </div>
            </Card>
          </Container>
          {/* Footer */}
          <Container
            maxWidth="md"
            component="footer"
            className={classes.footer}
          >
            <Box mt={5}>
              <Copyright />
            </Box>
          </Container>
        </ThemeProvider>
      )}
    </div>
  );
};

export default AskQuestion;
