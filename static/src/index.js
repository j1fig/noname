import 'bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import $ from 'jquery';
import _ from 'lodash';

function component() {
  let element = document.createElement('div');

  // element.innerHTML = _.join(['Hello', 'webpack'], ' ');
  $.ajax('api/stops');
  setTimeout(function () {
    console.log("hullo");
  }, 5000);

  return element;
}

document.body.appendChild(component());
