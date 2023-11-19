// // AccountOverviewForm.js
// import React, { useState } from 'react';

// function AccountOverviewForm() {
//   const [accountName, setAccountName] = useState('');
//   const [accountBalance, setAccountBalance] = useState('');

//   const handleSubmit = (e) => {
//     e.preventDefault();
//     // You can handle form submission logic here
//     console.log('Account Overview Form Submitted:', { accountName, accountBalance });
//   };

//   return (
//     <div>
//       <h2>Account Overview Form</h2>
//       <form onSubmit={handleSubmit}>
//         <label>
//           Account Name:
//           <input type="text" value={accountName} onChange={(e) => setAccountName(e.target.value)} />
//         </label>
//         <br />
//         <label>
//           Account Balance:
//           <input type="number" value={accountBalance} onChange={(e) => setAccountBalance(e.target.value)} />
//         </label>
//         <br />
//         <button type="submit">Submit</button>
//       </form>
//     </div>
//   );
// }

// export default AccountOverviewForm;
