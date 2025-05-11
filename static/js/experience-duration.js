function diffYearsMonths(from, to = new Date()) {
  let years  = to.getFullYear() - from.getFullYear();
  let months = to.getMonth()    - from.getMonth();
  if (to.getDate() < from.getDate()) months -= 1;      // partial month → not full
  if (months < 0) { years -= 1; months += 12; }         // borrow a year
  return { years, months };
}

document.addEventListener("DOMContentLoaded", () => {
  document.querySelectorAll(".experience-period").forEach(node => {
    const start = new Date(node.dataset.start);               // required
    const end   = node.dataset.end === "present"
                ? new Date()
                : new Date(node.dataset.end);

    const { years, months } = diffYearsMonths(start, end);

    // Build a human-readable string
    const ys = years  ? `${years} year${years  !== 1 ? "s" : ""}` : "";
    const ms = months ? `${months} month${months !== 1 ? "s" : ""}` : "";
    node.querySelector(".duration").textContent =
      ys && ms ? `${ys} ${ms}` : ys || ms || "Less than a month";
  });
});

