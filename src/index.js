function getSelectedUnits() {
    const currentUnit = document.querySelector('input[name="column1"]:checked');
    const convertUnit = document.querySelector('input[name="column2"]:checked');
    
    const selections = {
        current: currentUnit ? currentUnit.value : null,
        convert: convertUnit ? convertUnit.value : null
    };
    
    return selections;
}

function getScale(measure) {
    switch(measure) {
        case "millimeter":
            return 0.001
        case "centimeter":
            return 0.01
        case "decimeter":
            return 0.1
        case "meter":
            return 1.0
        case "dekameter":
            return 10.0
        case "hectometer":
            return 100.0
        case "kilometer":
            return 1000.0
        default:
            return -1
    }
}

function convertVal(val, current, convert) {
    val *= getScale(current)

    val /= getScale(convert)

    return val
}

let val;

document.getElementById("convert").onclick = function() {
    const val = document.getElementById("value-input").value;
    const selections = getSelectedUnits();

    result = convertVal(val, selections.current, selections.convert);

    document.getElementById("output-value").textContent = result + ' ' + selections.convert + 's'; 
}






/*
async function performConversion() {
    const valueInput = document.getElementById('value-input');
    const value = parseFloat(valueInput.value)
    if (isNaN(value)) {
        alert('Please enter a valid number');
        return;
    }
    
    const units = getSelectedUnits();
    
    if (!units.current || !units.convert) {
        alert('Please select both current and convert units');
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
        } else {
            alert(`Result: ${data.result} ${units.convert}`);
            // Or display it in a dedicated result element
        }
    } catch (error) {
        alert('Failed to connect to server: ' + error.message);
    }
}

// Attach the function to the convert button when the page loads
document.addEventListener('DOMContentLoaded', function() {
    const convertBtn = document.querySelector('.convert-btn');
    convertBtn.addEventListener('click', performConversion);
});
*/
