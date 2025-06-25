document.getElementById("reserve-form").onsubmit = async (e) => {
  e.preventDefault();
  const body = {
    user_id: parseInt(document.getElementById("user").value),
    table_id: parseInt(document.getElementById("table").value),
  };
  const res = await fetch("/reserve", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(body),
  });
  alert(await res.text());
};

document.getElementById("order-form").onsubmit = async (e) => {
  e.preventDefault();
  const body = {
    user_id: parseInt(document.getElementById("order-user").value),
    table_id: parseInt(document.getElementById("order-table").value),
    menu_item_ids: document
      .getElementById("items")
      .value.split(",")
      .map((id) => parseInt(id.trim())),
  };
  const res = await fetch("/order", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(body),
  });
  alert(await res.text());
};
