# Quiz Game

This project is a simple quiz game written in Go. The application reads questions and answers from a CSV file, presents the questions to the user, and checks the user's answers. The quiz is timed, and the user receives a score at the end.

## Features

- Read questions and answers from a CSV file.
- Present questions to the user one by one.
- Check user's answers and provide immediate feedback.
- Time limit for the entire quiz.
- Calculate and display the final score.

## Usage

### Running the Quiz Application

1. Build the Application:

   Ensure you have Go installed on your machine. Then, build the application using the following command:

   ```bash
   go build ./main.go

2. Run the Application:

   Use the following command to run the application. You can specify the CSV file and the time limit as flags:
  ./main.go -csv=path/to/quiz.csv -time=20

- csv: Path to the CSV file containing the quiz questions and answers. Each line should be in the format: question,answer.
- time: Time limit for the quiz in seconds (default is 20 seconds).

-Example of a csv file: 
	What is the capital of France?, Paris
	What is 2+2?,4
	Who wrote 'Hamlet'?, Shakespeare

 And to Run this : ./main.go -csv=quiz.csv -time=30

### Note

this project is foe personal purposes only and i did it with help of some recourses in order to learn 
i know there is a lot of space to improve and develop the code
i try my best to work on it in the future 


