Stock Application
1) Instance of a System
    - User
        AddUser()
        BuyStock()
    - Stock
        AddStock()
    - Profit
        Enable TimePeriodWise Metrics

2) User Entity
    Id
    Name
    - Stocks - [
        {
            Id
            Name
            number of buying
            current price
            avg price
            profit
                {
                    percentage
                    value
                    isNegative
                }
        }
    ]

3) Stock Entity
    - Id
    - Name
    - number of buyers
    - number of sellers
    - profit {
        percentage
        value
        isNegative
    }

4) Profit
    - Daywise
    - Monthly
    - Yearly
    - overall# machine-coding-practice
