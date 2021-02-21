import fs from 'fs';

export const linesCount = (fileName, onSuccess, onError) => {
    const processFile = (err, data) => {
        if (err) {
            onError('unable to open file ' + fileName);
        } else {
            onSuccess(data.toString().split('\n').length);
        }
    };

    fs.readFile(fileName, processFile);
};

export const linesCountP = (fileName) => {
    return new Promise((resolve, reject) => {
        fs.readFile(fileName, (err, data) => {
            if (err) {
                reject('unable to open file ' + fileName);
            } else {
                resolve(data.toString().split('\n').length);
            }
        });
    });
};