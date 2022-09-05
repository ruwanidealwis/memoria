window.onload = function() {
  document.getElementById("create-tab").addEventListener("click", function () {
    window.location = `/create`;
  });

  document.getElementById("logout-btn").addEventListener("click", function () {
    sessionStorage.setItem("userID", null);
    window.location = `/`;
  })

  loadUserScrapbooks();

  const newScrapbookForm = document.querySelector("#new-scrapbook-form");
  let createSbBtn = document.getElementById("submit-form-btn");

  createSbBtn.addEventListener("click", function (e) {
      submitForm(e, newScrapbookForm);
  });
};

loadUserScrapbooks = async (params) => {
  userID = sessionStorage.getItem("userID");
  res = await performHttpGetRequest(`/scrapbooks/${userID}`, params);
  console.log(res);
  for (let i=0; i<res.length; i++) {
    console.log(res[i]);
    displayNewScrapbook(res[i].name, res[i].id);
  }
}

async function submitForm(e, form) {
  e.preventDefault();
  // disable button while submitting
  let createSbBtn = document.getElementById("submit-form-btn");
  createSbBtn.disabled = true;
  setTimeout(() => (createSbBtn.disabled = false), 2000);

  let scrapbookName = document.getElementById("name").value;
  let userID = Number(sessionStorage.getItem("userID"));

  var jsonFormData = new Object();
  jsonFormData.name = scrapbookName;
  jsonFormData.user  = userID;

  console.log(jsonFormData);
  console.log(JSON.stringify(jsonFormData));

  const headers = buildHeaders();
  const response = await performJSONPostHttpRequest(
    `/scrapbook`,
    headers,
    jsonFormData
  );
  if (response) {
    sessionStorage.setItem("scrapbookID", response.ID);
    displayNewScrapbook(response.name, response.id);
    console.log(response);
  } else {
    alert(`An error occured submitting your form.`);
  }
}

function buildHeaders() {
  const headers = {
    "Content-Type": "application/json",
  };
  return headers;
}

async function performJSONPostHttpRequest(fetchLink, headers, body) {
  if (!fetchLink || !headers || !body) {
    throw new Error("Missing POST request parameters");
  }
  try {
    const rawResponse = await fetch(fetchLink, {
      method: "POST",
      headers: headers,
      body: JSON.stringify(body),
    });
    const content = await rawResponse.json();
    console.log(content);
    return content;
  } catch (err) {
    console.error(`Error at fetch POST: ${err}`);
    throw err;
  }
}

function displayNewScrapbook(name, id) {
  let newSb = document.createElement('div');
  newSb.id = `sp-${id}`;
  let sbTitle = document.createElement('div');
  newSb.classList.add('book');
  sbTitle.classList.add('book-title');
  let text = document.createTextNode(`${name}`);
  sbTitle.appendChild(text);
  newSb.appendChild(sbTitle);
  document.getElementById('book-container').appendChild(newSb);
}
