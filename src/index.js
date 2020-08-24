document.getElementById("app").innerHTML = `
<h1>Hello Vanilla!</h1>
<div>
  We use the same configuration as Parcel to bundle this sandbox, you can find more
  info about Parcel 
  <a href="https://parceljs.org" target="_blank" rel="noopener noreferrer">here</a>.
</div>
`;

const apiKey = `vsLelkOTHNWt7cHWX084zA==`;
const code = `DT_0001`;

let params = {
ServiceKey : `${apiKey}`,
ObsCode : `${code}`,
Date:"20200901",
ResultType:"json"
}
const url = `http://www.khoa.go.kr/oceangrid/grid/api/tideObsPreTab/search.do?ServiceKey=vsLelkOTHNWt7cHWX084zA==&ObsCode=DT_0001&Date=20200902`;
//http://www.khoa.go.kr/oceangrid/grid/api/tideObsPreTab/search.do?ServiceKey=vsLelkOTHNWt7cHWX084zA==&ObsCode=DT_0001&Date=20200902&ResultType=json
console.log(url);
function getData() {
  fetch(url)
    .then((response) => {
      return response.json();
    }).then(data=>{console.log(data)})
}


getData()
