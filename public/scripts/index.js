var d = document;
var users = [];
var user;
var fields = ['name', 'birth-date', 'address', 'phone', 'email'];
var inputElements = fields.map(function(field) {
  return d.querySelector(`[name="${field}"]`);
});

function SuperUser() {
  this.isDataVisible = true;
}

function hideUserData (e) {
  var el = e.target.parentElement;
  el.style.display = 'none';
  user.isDataVisible = false;
}

d.querySelector('tbody').addEventListener('click', hideUserData, false);

function User() {
  SuperUser.call(this);
  this.getFormValues = function () {
    var values = [];
    inputElements.forEach(function(el) {
      values.push(el.value);
    });
    var sex = d.querySelector('[name="sex"]:checked').value;
    values.push(sex);
    this.setUserData(values);
    users.push(this);
  };

  this.setUserData = function (values) {
    this.name = values[0];
    this.birth_date = values[1].split('-').reverse().join('.');
    this.address = values[2];
    this.phone = values[3];
    this.email = values[4];
    this.sex = values[5];
  };
}


function saveUser(e) {
  e.preventDefault();
  user = new User();
  user.setId();
  user.getFormValues();
  user.addTableRow();
  user.addCard();
  saveStudent(user.getUserData());
  clearInput();
  $submit.disabled = true;
}

User.prototype.setId = function () {
  this.id = Math.floor(Math.random() * 1000);
};

User.prototype.getUserData = function () {
  return [this.name, this.sex, this.birth_date, this.address, this.phone, this.email];
};


User.prototype.addTableRow = function () {
  var $tbody = d.querySelector('tbody');
  var createRow = d.createElement('tr');
  createRow.setAttribute('data-id', this.id);
  $tbody.appendChild(createRow);
  var userData = user.getUserData();
  var $tr = d.querySelector('[data-id=\"' + this.id + '\"]');
  userData.forEach(function (item) {
    var $tcell = d.createElement('td');
    $tcell.innerText = item;
    $tr.appendChild($tcell);
  });
};

User.prototype.addCard = function () {
  var $userList = d.querySelector('div.user-list');
  var $user = d.createElement('div');
  $user.setAttribute('data-id', this.id);
  $user.innerText = user.name;
  $userList.appendChild($user);
  d.querySelector('div[data-id=\"' + this.id + '\"]').addEventListener('click', showTableRow, false);
  };

function showTableRow (e) {
  var selectedUser = e.target.getAttribute('data-id');
  var $tr = d.querySelector('tr[data-id=\"' + selectedUser + '\"]');
  $tr.style.display = 'table-row';
}


function clearInput () {
  inputElements.forEach(function(el) {
    el.value = "";
  });
  d.querySelector('[value="Male"]').checked = true;
}

var $submit = d.querySelector('input[type="submit"]');
$submit.addEventListener('click', saveUser, false);

d.querySelector('form').addEventListener('change', submitEnable, false);

function submitEnable() {
  var validInput = 0;
  if (inputElements[0].value) {
    ( inputElements[0].value.match(/[a-zA-Z]+/)) ? (validInput++) : (alert('Name shouldn\'t contain any numbers or space'));
  }

  validInput = checkDate(inputElements[1], validInput);

  if (inputElements[3].value) {
    inputElements[3].value.match(/\+380\d{9}/) ? (validInput++) : (alert('Please, enter phone number using example'));
  }

  if (inputElements[4].value) {
    inputElements[4].value.match(/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/) ? (validInput++) : (alert('Please, enter email address using example'));
  }
  
  if (validInput === 4) {
    $submit.disabled = false;
  }
}

function checkDate(elem, validInput) {
  if (elem.value) {
    var date = elem.value.split('-');
    switch(date) {
      case (+date[0] < 1930):
        alert('Are you really that old?)');
        break;
      case (+date[1] > 12):
        alert('Check if you entered month correctly');
        break;
      case (+date[2] >= 31):
        alert('Check if you entered day of your birth correctly');
        break;
      default:
        return validInput +=1;
    }
  }
  return validInput;
}

function saveStudent(studentData) {
  studentData = {
    name: studentData[0],
    sex: studentData[1],
    birthDate: studentData[2],
    address: studentData[3],
    phone: studentData[4],
    email: studentData[5]
  }
  studentData = JSON.stringify(studentData);
  $.ajax({
    type: "POST",
    url: "/save_student",
    data: studentData,
    contentType:  'application/json'
  });
}