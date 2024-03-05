import fs from "fs";
import Handlebars from "./handlebars-wrapper.js";
import RpcParser from "./rpc-parser.js";

const main = (args) => {
  const [templatesDir, protoDir, destDir] = args;
  const parser = new RpcParser(protoDir);
  const services = parser.parse();

  const serviceTemplateText = fs.readFileSync(
    `${templatesDir}/client.go.hbs`,
    "utf-8"
  );
  const connectionTemplateText = fs.readFileSync(
      `${templatesDir}/connection.go.hbs`,
      "utf-8"
  );
  const connectionFooterTemplateText = fs.readFileSync(
    `${templatesDir}/connection_footer.go.hbs`,
    "utf-8"
  );
  const serviceTemplate = Handlebars.compile(serviceTemplateText);
  const connectionTemplate = Handlebars.compile(connectionTemplateText);
  const connectionFooterTemplate = Handlebars.compile(connectionFooterTemplateText);
  const serviceFileTemplate = Handlebars.compile('{{snakeCase serviceName}}_client.go');

  let footer = "";
  for (const service of services) {
    const serviceFileName = serviceFileTemplate(service);
    const serviceCodeGenerated = serviceTemplate(service);
    footer += connectionFooterTemplate(service);

    fs.writeFileSync(`${destDir}/${serviceFileName}`, serviceCodeGenerated);
  }

  fs.writeFileSync(`${destDir}/connection.go`, connectionTemplate({footer}));
};

main(process.argv.slice(2));
