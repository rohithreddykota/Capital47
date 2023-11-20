import React, { useState } from 'react';
import './CurrencySendingPage.css'; // Make sure to adjust the path based on your file structure

const CurrencySendingPage = () => {
    const [formData, setFormData] = useState({
        senderFirstName: '',
        senderLastName: '',
        recipientFirstName: '',
        recipientLastName: '',
        recipientCountry: 'usa',
        amount: '',
        currency: 'usd',
    });

    const handleInputChange = (e) => {
        const { name, value } = e.target;
        setFormData({
            ...formData,
            [name]: value,
        });
    };

    const submitForm = () => {
        alert('Form submitted!');
    };

    return (
        <div className="currency-sending-container">
            <h2 className="currency-sending-title">Send Money Overseas</h2>
            <form className="currency-sending-form">
                {/* Sender Information */}
                <label className="currency-sending-label" htmlFor="senderFirstName">Sender's First Name:</label>
                <input
                    className="currency-sending-input"
                    type="text"
                    id="senderFirstName"
                    name="senderFirstName"
                    value={formData.senderFirstName}
                    onChange={handleInputChange}
                    required
                />

                {/* Recipient Information */}
                <label className="currency-sending-label" htmlFor="recipientCountry">Recipient's Country:</label>
                <select
                    className="currency-sending-select"
                    id="recipientCountry"
                    name="recipientCountry"
                    value={formData.recipientCountry}
                    onChange={handleInputChange}
                    required
                >
                    <option value="usa">🇺🇸 United States</option>
                    <option value="uk">🇬🇧 United Kingdom</option>
                    <option value="india">🇮🇳 India</option>
                    <option value="australia">🇦🇺 Australia</option>
                    <option value="china">🇨🇳 China</option>
                    <option value="russia">🇷🇺 Russia</option>
                    <option value="italy">🇮🇹 Italy</option>
                    <option value="netherlands">🇳🇱 Netherlands</option>
                    <option value="canada">🇨🇦 Canada</option>
                    <option value="germany">🇩🇪 Germany</option>
                </select>

                {/* Amount and Currency */}
                <label className="currency-sending-label" htmlFor="amount">Amount to Send:</label>
                <input
                    className="currency-sending-input"
                    type="number"
                    id="amount"
                    name="amount"
                    value={formData.amount}
                    onChange={handleInputChange}
                    required
                />

                <label className="currency-sending-label" htmlFor="currency">Currency:</label>
                <select
                    className="currency-sending-select"
                    id="currency"
                    name="currency"
                    value={formData.currency}
                    onChange={handleInputChange}
                    required
                >
                    <option value="usd">🇺🇸 USD</option>
                    <option value="eur">🇪🇺 EUR</option>
                    <option value="aud">🇦🇺 AUD</option>
                    <option value="inr">🇮🇳 INR</option>
                    <option value="cyp">🇨🇾 CYP</option>
                    <option value="bsd">🇧🇸 BSD</option>
                    <option value="bgn">🇧🇬 BGN</option>
                    <option value="cad">🇨🇦 CAD</option>
                    <option value="jpy">🇯🇵 JPY</option>
                    <option value="gbp">🇬🇧 GBP</option>
                </select>

                {/* Submit Button */}
                <button className="currency-sending-button" type="button" onClick={submitForm}>
                    Send Money
                </button>
            </form>
        </div>
    );
};

export default CurrencySendingPage;
