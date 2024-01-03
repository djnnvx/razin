rule RAZIN_REVSHELL {
    meta:
        description = "Detects golang reverse-shell implant"
        author = "djnn"
        date = "2024-04-01"
        reference = "https://github.com/djnnvx/razin"
        hash1 = "eb5a9b59f279a0c6d552847e57baf794d4cb67e37d369bee8245855d2a737839"

    strings:
        $str1 = "RAZINrazinRAZINrazinRAZINrazinRAZINraz"
        $str2 = "10.0.2.2"
        $str3 = "github.com/bogdzn/razin"

    condition:
        3 of them

}
