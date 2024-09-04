<template>
    <body>
      <div class = "container" >
        
        <input class = "textBox" placeholder="Username">
        </input>

        <input class = "textBox" placeholder="Password">
        </input>


        <button class = "buttons">
            <slot>
                Sign Up
            </slot>
        </button>
      </div>
    </body>
  </template>