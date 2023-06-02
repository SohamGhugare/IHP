import { Button,Typography, FormControl, InputLabel, MenuItem, Select, TextField } from '@mui/material';

const MyForm = () => {
  const handleSubmit = (e) => {
    e.preventDefault();
    // Handle form submission
  };

  return (

    <div className="flex items-center justify-center ">
      <form onSubmit={handleSubmit} className="w-full max-w-md">
        <Typography variant="h2" className="text-2xl font-bold mb-6">
          Login
        </Typography>

        <TextField
          id="fullName"
          label="Full Name"
          variant="outlined"
          margin="normal"
          fullWidth
          required
          InputLabelProps={{
            style: { color: '#838383' },
          }}
        />

        <TextField
          id="phoneNumber"
          label="Phone Number"
          variant="outlined"
          margin="normal"
          fullWidth
          required
          InputLabelProps={{
            style: { color: '#838383' },
          }}
        />

        {/* Rest of the input fields with customized label color */}

        <Button
          type="submit"
          variant="contained"
          color="primary"
          fullWidth
        >
          Login
        </Button>
      </form>
    </div>



  );
};

export default MyForm;
