import React, { useEffect, useState } from "react";
import CssBaseline from "@material-ui/core/CssBaseline";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";
import Link from "@material-ui/core/Link";
import Container from "@material-ui/core/Container";
import Box from "@material-ui/core/Box";
import Header from "../components/Header";
import Questions from "../components/Questions";
import { useDispatch, useSelector } from "react-redux";
import { CircularProgress } from "@material-ui/core";
import questionAction from "../redux/question/questionAction";
import useStyles from "../styles/Style";
import Copyright from "../components/Copyright";
import categoryAction from "../redux/category/get/categoryAction";

const NotFound = () => {
  const classes = useStyles();

  return (
    <React.Fragment>
      <>
        <CssBaseline />
        <Header />
        {/* Hero unit */}
        <Container
          maxWidth="md"
          component="main"
          className={classes.heroContent}
        >
          <h1>404 Page Not Found</h1>
        </Container>
        {/* End hero unit */}
        <Container maxWidth="md" component="main">
          <Grid container spacing={5} alignItems="flex-end"></Grid>
        </Container>
        {/* Footer */}
        <Container maxWidth="md" component="footer" className={classes.footer}>
          <Box mt={5}>
            <Copyright />
          </Box>
        </Container>
      </>
      {/* End footer */}
    </React.Fragment>
  );
};

export default NotFound;
