package config

type Resolution string

type Distribution struct {
    Name string
}

type User struct {
    Name string
}

type Box struct {
    Distribution Distribution
    User User
}

type Environment struct {
    Resolution Resolution
    Var []string
}

type Command struct {
    Environment Environment
}

type Config struct {
    Box Box
    Command Command
}
