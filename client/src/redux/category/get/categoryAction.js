import stuckoverflowAPI from "../../../APIs/stuckoverflow";

const fetchCategoryList = () => async (dispatch) => {
  try {
    const categoryList = await stuckoverflowAPI({
      method: "GET",
      url: `/categories`,
    });

    dispatch({
      type: "SEARCH_CATEGORY",
      payload: {
        list: categoryList.data.data
      }
    });

  } catch (error) {
    console.log(error);
  }
};

const categoryAction = {
  fetchCategoryList,
};

export default categoryAction;