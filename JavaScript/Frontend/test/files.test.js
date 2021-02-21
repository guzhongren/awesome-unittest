import { linesCount, linesCountP } from '../src/files';

describe('test server side callback', () => {
    it('should return correct lines count for a valid file', (done) => {
        const onSuccess = (data) => {
            try {
                expect(data).toBe(25);
                done();
            } catch (error) {
                done(error);
            }
        };

        linesCount('src/files.js', onSuccess);
    });

    it('should report error for an invalid file name', (done) => {
        const onError = (err) => {
            expect(err).toBe('unable to open file src/async/files.js');
            done();
        };

        linesCount('src/async/files.js', undefined, onError);
    });
});

describe('test server side promise', () => {
    // 使用 done 和 promise
    it('should return correct lines count for valid file - with done', (done) => {
       const checkCount = (count) => {
           expect(count).toBe(25);
           done();
       }; 

       linesCountP('src/files.js')
                    .then(checkCount);
    });

    // 使用promise
    it('should return correct lines count for valid file - return promise', () => {
        // const checkCount = (count) => {
        //     expect(count).toBe(25);
        // }; 

        // return linesCountP('src/files.js')
        //                 .then(checkCount);
        // 简化
        return expect(linesCountP('src/files.js')).resolves.toBe(25);
    });

    it('should report error for an invalid file name - return promise', () => {
        // const onError = (err) => {
        //     expect(err).toBe('unable to open file src/async/files.js');
        // }; 
        // expect.assertions(1);
        // return linesCountP('src/async/files.js').catch(onError);
        // 简化
        return expect(linesCountP('src/async/files.js')).rejects.toMatch('unable to open file src/async/files.js');
    });

    // 使用async await
    it('should return correct lines count for valid file - with async', async () => {
        const data = await linesCountP('src/files.js');
        expect(data).toBe(25);
    });

    it('should report error for an invalid file name - with async', async () => {
        expect.assertions(1);
        try {
            await linesCountP('src/async/files.js');
        } catch (error) {
            expect(error).toMatch('unable to open file src/async/files.js');
        }
    });

    // 结合使用 promise 和 async await
    it('should return correct lines count for valid file - with promise and async', async () => {
        await expect(linesCountP('src/files.js')).resolves.toBe(25);
    });

    it('should report error for an invalid file name - with promise and async', async () => {
        await expect(linesCountP('src/async/files.js')).rejects.toMatch('unable to open file src/async/files.js');
    });
});
