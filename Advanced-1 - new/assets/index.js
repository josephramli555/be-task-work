$(document).ready(function () {
  $(".delete-item").on("click", function (e) {
    let url = "http://localhost:8080/answers/" + e.target.id;
    fetch(url, {
      method: "DELETE",
    }).then(function (res) {
      return (window.location.href = "/report");
    });
  });

  $("#form-update").submit(function (e) {
    e.preventDefault();
    let updateId = $("#updateId").val();
    let newAnswer = $("#updateAnswer").val();
    let url = "http://localhost:8080/answers/" + updateId;
    let formData = new FormData();
    formData.append("answer", newAnswer);
    fetch(url, {
      body: formData,
      method: "PUT",
    }).then(function (res) {
      return (window.location.href = "/report");
    });
  });
});
