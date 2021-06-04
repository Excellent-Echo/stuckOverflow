import axios from "axios";
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
  ThemeProvider,
  withStyles,
} from "@material-ui/core";
import Theme from "../styles/Theme";
import useStyles from "../styles/Style";
import HeaderAppBar from "./Header";

URL = "https://stuckoverflow.herokuapp.com/questions"; // URL variable stores JSON url || API taken from 10 Degrees WordPress Agency

class Search extends React.Component {
  state = {
    post: [],
    allPosts: [],
  };

  componentDidMount() {
    axios
      .get(URL, {
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
      })
      .then(({ data }) => {
        this.setState({
          post: data.data,
          allPosts: data.data, // array data from JSON stored in these
        });
      })
      .catch((err) => {});
  }

  _onKeyUp = (e) => {
    // filter post list by title using onKeyUp function
    const post = this.state.allPosts.filter((item) =>
      item.title.toLowerCase().includes(e.target.value.toLowerCase())
    );
    this.setState({ post });
  };

  render() {
    const { classes } = this.props;
    return (
      <>
        <ThemeProvider theme={Theme}>
          <HeaderAppBar />
          <Container
            maxWidth="md"
            component="main"
            className={classes.heroContent}
          >
            <Grid item xs={12} md={12}>
              <Box display="flex" justifyContent="space-between">
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
            </Grid>
          </Container>
          <div className="container">
            <div className="search-outer">
              <form
                role="search"
                method="get"
                id="searchform"
                className="searchform"
                action=""
              >
                {/* input field activates onKeyUp function on state change */}
                <input
                  type="search"
                  onChange={this._onKeyUp}
                  name="s"
                  id="s"
                  placeholder="Search"
                />
                <button type="submit" id="searchsubmit">
                  <i className="fa fa-search" aria-hidden="true" />
                </button>
              </form>
            </div>
            <ul className="data-list">
              {/* post items mapped in a list linked to onKeyUp function */}
              {this.state.post.map((item, index) => (
                <li className={"block-" + index}>
                  <a className="title" href={item.link}>
                    <h3>{item.title}</h3>
                  </a>
                  <a className="link" href={item.link}></a>
                </li>
              ))}
            </ul>
          </div>
        </ThemeProvider>
      </>
    );
  }
}

export default withStyles(useStyles)(Search);
