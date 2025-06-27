# The Drunken Jailer

## Scenario
> In a prison there are `N` cells arranged in a row. One day the jailer got drunk. He run thru the prison `N` times, and in the `I`<sup>th</sup> time he toggle the state of the door of every `I`<sup>th</sup> cell.

> That is, for the `1`<sup>st</sup> time he unlocked all the cell doors. In the `2`<sup>nd</sup> time he locked back the doors of cell 2, 4, 6, 8... Then in the `3`<sup>rd</sup> time he went to cell 3, 6, 9, 12... and relocked the unlocked doors, and unlocked the locked doors. The jailer repeated the process `N` times before passing out.

> For a given integer of `N`, determine the number of cell doors remain opened after the process.

## Hints
1. Direct simulation is possible.

1. A cell door is unlock if the cell's number has an odd number of divisors.

1. A perfect square number (e.g. 1, 4, 9, 16...) has odd number divisors.
