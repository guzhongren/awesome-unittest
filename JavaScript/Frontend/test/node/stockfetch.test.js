// 选择有趣且重要的函数，为其首先编写测试，也是一个好的起点 - read函数
describe('Stockfetch test', () => {
  it('should pass this canary test', () => {
    expect(true).toBe(true);
  });

  it('read should invoke error handler for invalid file', (done) => {
    const onError = (err) => {
      expect(err).toEqual('Error reading file: InvalidFile');
      done();
    };

    
  });
});
