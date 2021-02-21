// 原型代码
// function locate() {
//   console.log(1);
//   navigator.geolocation.getCurrentPosition(
//     function (position) {
//       const latitude = position.coords.latitude;
//       const longitude = position.coords.longitude;

//       const url = `http://maps.google.com/?q=${latitude},${longitude}`;
//       window.location = url;
//     },
//     function () {
//       document.getElementById("error").innerHTML =
//         "unable to get your location";
//     }
//   );
// }

export const createUrl = (lat, lon) => {
  if (!lat || !lon) {
    return "";
  }
  return `https://map.baidu.com/poi/@${lat},${lon}`;
}; 

export const setLocation = (windowObj, url) => {
    windowObj.location = url;
}

export const onSuccess = (position) => {};

export const onError = (error) => {
    document.getElementById('error').innerHTML = error.message;
};

export const locate = () => {
    navigator.geolocation.getCurrentPosition(onSuccess, onError)
};
