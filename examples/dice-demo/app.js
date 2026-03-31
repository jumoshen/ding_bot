const dice = document.getElementById("dice");
const diceWrap = document.getElementById("dice-wrap");
const shadow = document.getElementById("dice-shadow");
const rollBtn = document.getElementById("roll-btn");
const resultText = document.getElementById("result");
const targetSelect = document.getElementById("target");
const speedInput = document.getElementById("speed");
const speedValue = document.getElementById("speed-value");

const FACE_ROTATION = {
  1: { x: 0, y: 0 }, // front
  2: { x: 90, y: 0 }, // bottom -> front
  3: { x: 0, y: -90 }, // right -> front
  4: { x: 0, y: 90 }, // left -> front
  5: { x: -90, y: 0 }, // top -> front
  6: { x: 180, y: 0 }, // back -> front
};

const state = {
  x: -24,
  y: 36,
  z: -12,
  rolling: false,
};

function setDiceRotation(x, y, z = state.z) {
  state.x = x;
  state.y = y;
  state.z = z;
  dice.style.transform = `rotateX(${x}deg) rotateY(${y}deg) rotateZ(${z}deg)`;
}

function getRandomInt(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

function randomChoice(list) {
  return list[getRandomInt(0, list.length - 1)];
}

function updateSpeedLabel() {
  speedValue.textContent = `${Number(speedInput.value).toFixed(2)}x`;
}

function setRollingUI(isRolling) {
  state.rolling = isRolling;
  rollBtn.disabled = isRolling;
  rollBtn.textContent = isRolling ? "转动中..." : "掷骰子";
}

function animateBounce(duration) {
  diceWrap.animate(
    [
      { transform: "translateY(-20px)" },
      { transform: "translateY(0px)" },
      { transform: "translateY(-7px)" },
      { transform: "translateY(0px)" },
    ],
    {
      duration,
      easing: "cubic-bezier(.16,.84,.23,1)",
      fill: "forwards",
    }
  );

  shadow.animate(
    [
      { transform: "scale(0.75)", opacity: 0.22 },
      { transform: "scale(1)", opacity: 0.3 },
      { transform: "scale(0.92)", opacity: 0.26 },
      { transform: "scale(1)", opacity: 0.3 },
    ],
    {
      duration,
      easing: "cubic-bezier(.16,.84,.23,1)",
      fill: "forwards",
    }
  );
}

function animatePreLift(duration) {
  diceWrap.animate(
    [{ transform: "translateY(0px)" }, { transform: "translateY(-20px)" }],
    {
      duration,
      easing: "cubic-bezier(.2,.8,.3,1)",
      fill: "forwards",
    }
  );

  shadow.animate(
    [{ transform: "scale(1)", opacity: 0.3 }, { transform: "scale(0.72)", opacity: 0.22 }],
    {
      duration,
      easing: "cubic-bezier(.2,.8,.3,1)",
      fill: "forwards",
    }
  );
}

function rollDice(targetValue) {
  if (state.rolling) {
    return;
  }
  setRollingUI(true);

  const result = targetValue ?? getRandomInt(1, 6);
  const speed = Number(speedInput.value);
  const baseFace = FACE_ROTATION[result];

  const spinX = randomChoice([720, 900, 1080, 1260, 1440]);
  const spinY = randomChoice([720, 900, 1080, 1260, 1440]);
  const spinZ = randomChoice([180, 360, 540, 720]);
  const tilt = randomChoice([-12, -8, -4, 4, 8, 12]);

  const targetX = baseFace.x + spinX;
  const targetY = baseFace.y + spinY;
  const targetZ = spinZ + tilt;

  const preDuration = Math.round(110 / speed);
  const spinDuration = Math.round(760 / speed);
  const settleDuration = Math.round(240 / speed);

  animatePreLift(preDuration);

  const preKeyframes = [
    { transform: `rotateX(${state.x}deg) rotateY(${state.y}deg) rotateZ(${state.z}deg)` },
    { transform: `rotateX(${state.x + 16}deg) rotateY(${state.y - 14}deg) rotateZ(${state.z + 8}deg)` },
  ];

  const spinKeyframes = [
    { transform: preKeyframes[1].transform, filter: "blur(0px)" },
    { transform: `rotateX(${targetX * 0.55}deg) rotateY(${targetY * 0.52}deg) rotateZ(${targetZ * 0.6}deg)`, filter: "blur(0.4px)" },
    { transform: `rotateX(${targetX * 0.88}deg) rotateY(${targetY * 0.9}deg) rotateZ(${targetZ * 0.94}deg)`, filter: "blur(0.2px)" },
    { transform: `rotateX(${targetX}deg) rotateY(${targetY}deg) rotateZ(${targetZ}deg)`, filter: "blur(0px)" },
  ];

  const preAnim = dice.animate(preKeyframes, {
    duration: preDuration,
    easing: "cubic-bezier(.18,.78,.34,.99)",
    fill: "forwards",
  });

  preAnim.onfinish = () => {
    const spinAnim = dice.animate(spinKeyframes, {
      duration: spinDuration,
      easing: "cubic-bezier(.12,.83,.19,.99)",
      fill: "forwards",
    });

    spinAnim.onfinish = () => {
      const finalX = baseFace.x + 360 * 4;
      const finalY = baseFace.y + 360 * 4;
      const finalZ = randomChoice([-3, 0, 3]);

      const settleAnim = dice.animate(
        [
          { transform: `rotateX(${targetX}deg) rotateY(${targetY}deg) rotateZ(${targetZ}deg)` },
          { transform: `rotateX(${finalX + 8}deg) rotateY(${finalY - 6}deg) rotateZ(${finalZ + 3}deg)` },
          { transform: `rotateX(${finalX}deg) rotateY(${finalY}deg) rotateZ(${finalZ}deg)` },
        ],
        {
          duration: settleDuration,
          easing: "cubic-bezier(.2,.95,.3,1)",
          fill: "forwards",
        }
      );

      animateBounce(settleDuration);

      settleAnim.onfinish = () => {
        setDiceRotation(finalX, finalY, finalZ);
        resultText.textContent = String(result);
        setRollingUI(false);
      };
    };
  };
}

speedInput.addEventListener("input", updateSpeedLabel);
rollBtn.addEventListener("click", () => {
  const selected = targetSelect.value;
  const target = selected === "random" ? undefined : Number(selected);
  rollDice(target);
});

updateSpeedLabel();
setDiceRotation(state.x, state.y, state.z);
