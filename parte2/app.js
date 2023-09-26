const fs = require('fs');
const csv = require('csv-parser');

const inputFile = 'datos.csv';
const outputFile = 'datos.json';

const data = {};

fs.createReadStream(inputFile)
  .pipe(csv())
  .on('data', (row) => {

    const organization = row.organizacion;
    const username = row.usuario;
    const role = row.rol;

    if (!data[organization]) {
      data[organization] = {};
    }

    if (!data[organization][username]) {
      data[organization][username] = { roles: [] };
    }

    data[organization][username].roles.push(role);
  })
  .on('end', () => {
    const result = [];

    for (const organization in data) {
      const users = [];

      for (const username in data[organization]) {
        users.push({
          username: username,
          roles: data[organization][username].roles,
        });
      }

      result.push({
        organization: organization,
        users: users,
      });
    }

    fs.writeFileSync(outputFile, JSON.stringify(result, null, 2));
    console.log(`El archivo JSON "${outputFile}" ha sido creado exitosamente.`);
  });
