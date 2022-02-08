// window.addEventListener("load", () => {
const getCount = async () => {
  try {
    const response = await fetch("/counter");
    return await response.text();
  } catch (error) {
    console.log(error);
  }
};
const renderCount = async () => {
  let count = await getCount();
  const counterDisplay = document.getElementById("counter");
  counterDisplay.innerHTML = `This document has been viewed ${count} times.`;
};
renderCount();
// });
