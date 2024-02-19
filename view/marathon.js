"use strict";
var _a;
console.log('linked');
function generateTrainingSchedule(startDate, durationDays, startingKilometer, startingPace, targetPace, daysOfTraining) {
    const daysPerWeek = 2;
    const trainingDays = (durationDays / 7) * daysOfTraining;
    const weeklyIncrease = (startingKilometer / durationDays) * daysPerWeek;
    const test = (42.2 - startingKilometer) / trainingDays;
    const schedule = [];
    let currentDate = new Date(startDate);
    let currentDistance = startingKilometer;
    let currentPace = startingPace;
    for (let day = 0; day < trainingDays; day++) {
        schedule.push({
            distance: parseFloat(currentDistance.toFixed(2)),
            pace: parseFloat(currentPace.toFixed(2)),
        });
        if (currentDistance >= 42.2) {
            break;
        }
        currentDate.setDate(currentDate.getDate() + 1);
        currentDistance += weeklyIncrease;
        currentPace = targetPace;
    }
    return schedule;
}
//input date for when the marathon is
const marathonDateInput = document.getElementById("marathonDate");
const marathonDate = marathonDateInput ? new Date(marathonDateInput.value) : new Date();
//input date for when you want to start training
const startingTrainingInput = (_a = document.getElementById("date")) === null || _a === void 0 ? void 0 : _a.value; // Replace with user input (16 weeks)
const dateStartingTraining = startingTrainingInput ? new Date(startingTrainingInput) : new Date();
// how many kilometers you can run
const startingKilometer = document.getElementById("startingDistance"); // Replace with user input (10 kilometers)
// the pace that you can run now
const startingPace = document.getElementById("startingSpeed"); // Replace with user input (6 minutes per kilometer)
// the desirec pace that you want to run
const desiredSpeed = document.getElementById("desiredSpeed");
//days of training
const daysOfTraining = document.getElementById("daysOfTraining");
/// caluculatebutton
const calculateButton = document.getElementById("calculateButton");
let diffrenceInMiliseconds = (marathonDate.getTime() - dateStartingTraining.getTime());
const diffrenceInDays = diffrenceInMiliseconds / (1000 * 3600 * 24);
console.log(diffrenceInDays, 'diffrence in time');
console.log(marathonDate, 'start date', dateStartingTraining, 'date that you wanna start training');
let startingSpeedOutput = document.getElementById("startingSpeedOutput");
let desiredSpeedOutput = document.getElementById("desiredSpeedOutput");
// Update the current slider value (each time you drag the slider handle)
startingPace.oninput = function () {
    startingSpeedOutput.textContent = startingPace.value;
};
desiredSpeed.oninput = function () {
    desiredSpeedOutput.textContent = desiredSpeed.value;
};
const warningMessage = document.getElementById("message");
console.log(startingPace.value);
calculateButton.addEventListener('click', () => {
    if (startingPace.value === '' || desiredSpeed.value === '' || startingKilometer.value === '') {
        warningMessage.textContent = 'Please fill in all the fields';
    }
    else {
        warningMessage.textContent = '';
        const trainingSchedule = generateTrainingSchedule(dateStartingTraining, diffrenceInDays, parseInt(startingKilometer.value), parseFloat(startingPace.value), parseFloat(desiredSpeed.value), parseInt(daysOfTraining.value));
        console.log(trainingSchedule);
    }
});
