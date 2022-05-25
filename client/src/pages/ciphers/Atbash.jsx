import { Button, Divider, Typography } from "@mui/material";
import { TextField } from "@mui/material";

export default function Atbash(props) {
    return <>
        <Typography variant="h2" gutterBottom >Atbash</Typography>

        <Typography variant="h4" gutterBottom>Morse is an encoding method that maps every ASCII letter to a code consisting of "." and "-".</Typography>

        <Typography variant="body1">To increase the efficiency of encoding, Morse code was designed so that the length of each symbol is approximately inverse to the frequency of occurrence of the character that it represents in text of the English language. Thus the most common letter in English, the letter e, has the shortest code: a single dit. Because the Morse code elements are specified by proportion rather than specific time durations, the code is usually transmitted at the highest rate that the receiver is capable of decoding. Morse code transmission rate (speed) is specified in groups per minute, commonly referred to as words per minute</Typography>

        <Divider sx={{ margin: "20px 0" }} />



        <TextField
            id="text"
            label="Text"
            placeholder="Text to process"
            multiline
            rows={4}
            required
            value={props.data["text"]}
            onChange={event => props.changeHandler("text", event.target.value)}
            fullWidth
            sx={{
                marginBottom: "20px",
            }}
        />

        <Button variant="contained" onClick={props.submitHandler} sx={{
            width: "300px",
            height: "50px",
            display: "block",
            margin: "auto",
            marginBottom: "20px"

        }}>Go!</Button>

    </>
}