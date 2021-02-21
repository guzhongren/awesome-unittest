function randomUser() {
  const userDom =  document.getElementById('phone');
  axios({
    method:'get',
    url:'https://randomuser.me/api/'
  })
  .then(function (response) {
    
  })
  .catch(function (error) {
    // ?.innerHTML = '暂无图像等信息';
  });
}