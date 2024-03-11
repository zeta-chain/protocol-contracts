module.exports.names = (params) => params.map((p) => p.name).join(", ");

module.exports["typed-params"] = (params) => {
  return params?.map((p) => `${p.type}${p.indexed ? " indexed" : ""}${p.name ? " " + p.name : ""}`).join(", ");
};

const slug = (module.exports.slug = (str) => {
  if (str === undefined) {
    throw new Error("Missing argument");
  }
  return str.replace(/\W/g, "-");
});
