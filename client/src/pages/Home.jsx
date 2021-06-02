import React, { useEffect } from "react";
import CssBaseline from "@material-ui/core/CssBaseline";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";
import Link from "@material-ui/core/Link";
import { makeStyles } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";
import Box from "@material-ui/core/Box";
import Header from "../components/Header";
import Questions from "../components/Questions";
import { useDispatch, useSelector } from "react-redux";
import { CircularProgress } from "@material-ui/core";
import questionAction from "../redux/question/questionAction";
import useStyles from "../styles/Style";

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {"Copyright © "}
      <Link color="inherit" href="https://material-ui.com/">
        Stuck Overflow
      </Link>{" "}
      {new Date().getFullYear()}
    </Typography>
  );
}

const Home = () => {
  const classes = useStyles();
  const questionsData = useSelector((state) => state.question);
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(questionAction.fetchAllQuestions());
  }, []);

  return (
    <React.Fragment>
      {Object.keys(questionsData).length === 0 ? (
        <Box
          display="flex"
          justifyContent="center"
          alignItems="center"
          minHeight="100vh"
        >
          <CircularProgress />
        </Box>
      ) : (
        <>
          <CssBaseline />
          <Header />
          {/* Hero unit */}
          <Container
            maxWidth="md"
            component="main"
            className={classes.heroContent}
          >
            <Questions />
          </Container>
          {/* End hero unit */}
          <Container maxWidth="md" component="main">
            <Grid container spacing={5} alignItems="flex-end"></Grid>
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
        </>
      )}
      {/* End footer */}
    </React.Fragment>
  );
};

export default Home;
