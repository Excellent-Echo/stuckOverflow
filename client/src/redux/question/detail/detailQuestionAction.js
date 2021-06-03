import stuckoverflowAPI from "../../../APIs/stuckoverflow";

const fetchDetailQuestion = (id) => async (dispatch) => {
  try {
    const questions = await stuckoverflowAPI({
      method: "GET",
      url: `/questions/${id}`,
    });

    dispatch({
      type: "SHOW_QUESTION",
      payload: {
        data: questions.data,
      }
    });

  } catch (error) {
    console.log(error);
  }
};

const removeDetailQuestion = () => {
  return {
    type: "REMOVE_DETAIL",
  }
};

const detailQuestionAction = {
  fetchDetailQuestion,
  removeDetailQuestion,
};

export default detailQuestionAction;