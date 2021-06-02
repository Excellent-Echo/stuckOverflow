import stuckoverflowAPI from "../../APIs/stuckoverflow";

const fetchAllQuestions = () => async (dispatch) => {
  try {
    const questions = await stuckoverflowAPI({
      method: "GET",
      url: "/questions",
    });

    dispatch({
      type: "SHOW_ALL_QUESTIONS",
      payload: {
        loaded: false,
        list: questions.data.data,
      }
    });

  } catch (error) {
    console.log(error);
  }
};

const questionAction = {
  fetchAllQuestions,
};

export default questionAction;