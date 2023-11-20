// // TransactionDetailsForm.js
// import React, { useState } from 'react';

// function TransactionDetailsForm() {
//   const [transactionType, setTransactionType] = useState('');
//   const [amount, setAmount] = useState('');

//   const handleSubmit = (e) => {
//     e.preventDefault();
//     // You can handle form submission logic here
//     console.log('Transaction Details Form Submitted:', { transactionType, amount });
//   };

//   return (
//     <div>
//       <h2>Transaction Details Form</h2>
//       <form onSubmit={handleSubmit}>
//         <label>
//           Transaction Type:
//           <input type="text" value={transactionType} onChange={(e) => setTransactionType(e.target.value)} />
//         </label>
//         <br />
//         <label>
//           Amount:
//           <input type="number" value={amount} onChange={(e) => setAmount(e.target.value)} />
//         </label>
//         <br />
//         <button type="submit">Submit</button>
//       </form>
//     </div>
//   );
// }

// export default TransactionDetailsForm;
