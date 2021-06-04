import stuckoverflowAPI from "../../../APIs/stuckoverflow";

const fetchSearchQuestions = (query) => async (dispatch) => {
  try {
    const questions = await stuckoverflowAPI({
      method: "GET",
      url: `/questions/?search=${query}`,
    });

    dispatch({
      type: "SEARCH_QUESTIONS",
      payload: {
        loaded: false,
        list: questions.data.data,
      }
    });

  } catch (error) {
    console.log(error);
  }
};

const searchQuestionAction = {
  fetchSearchQuestions,
};

export default searchQuestionAction;