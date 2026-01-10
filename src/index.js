function getSelectedUnits() {
    const currentUnit = document.querySelector('input[name="column1"]:checked');
    const convertUnit = document.querySelector('input[name="column2"]:checked');
    
    const selections = {
        current: currentUnit ? currentUnit.value : null,
        convert: convertUnit ? convertUnit.value : null
    };
    
    return selections;
}

async function performConversion() {
    const valueInput = document.getElementById('value-input');
    const value = parseFloat(valueInput.value);
    const outputElement = document.getElementById('output-value');
    
    if (isNaN(value)) {
        alert('Please enter a valid number');
        outputElement.textContent = '';
        return;
    }
    
    const units = getSelectedUnits();
    
    if (!units.current || !units.convert) {
        alert('Please select both current and convert units');
        outputElement.textContent = '';
        return;
    }
    
    try {
        const response = await fetch('http://localhost:8080/convert', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                value: value,
                current: units.current,
                convert: units.convert
            })
        });
        
        const data = await response.json();
        
        if (data.error) {
            alert('Error: ' + data.error);

            outputElement.textContent = '';
        } else {
            outputElement.textContent = data.result + ' ' + units.convert;
        }
    } catch (error) {
        alert('Failed to connect to server: ' + error.message);
        outputElement.textContent = '';
    }
}

// Attach the function to the convert button when the page loads
document.addEventListener('DOMContentLoaded', function() {
    const convertBtn = document.querySelector('.convert-btn');
    convertBtn.addEventListener('click', performConversion);
});
