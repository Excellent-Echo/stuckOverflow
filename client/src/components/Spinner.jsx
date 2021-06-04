import React from "react";

const LoadingSpinner = () => {
  return (
    <div>
      <div
        className="d-flex justify-content-center"
        style={{ marginTop: "200px" }}
      >
        <div className="spinner-border" role="status">
          <span className="visually-hidden"></span>
        </div>
      </div>
    </div>
  );
};

export default LoadingSpinner;
