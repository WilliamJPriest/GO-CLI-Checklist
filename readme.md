<p align="center">
  <img width="70%"  src="https://i.ibb.co/272d4rr/fc28f9f73dea4599b2d5ac2b3cacc13e.jpg"/>
</p>


# **Go Cli Checklist**


Welcome to Go CLI Checklist, this is a CLI / Tui based Task manager. Everything is stored locally so there's no need to be online or fear forgetting your password.

## **Basic Setup**
---

Make sure you have GO installed and then follow these steps:

1. Git clone https://github.com/WilliamJPriest/GO-CLI-Checklist

2. in the cloned directory type go build

3. After the build phase type go install

4. to run program type checklist into your terminal

5. Ta Da, everything should be working.

If an issue with install arises, please check the docs [here](https://go.dev/doc/tutorial/compile-install) 

### **Note:**

The CSV is stored on your user directory if you ever need to locate it.

i.e

`C:\Users\Lenovo\checklists.csv` 

## **How to Use**
---

After you run checklist you should see
```
    Select Action: 

	Create
	Read
	Update
	Delete
	Millionaire

```

Type the name of action you wish to use, it is not case-sensitive

i.e read

```
    read

    id: hiWorld ☑ Make a hello world program for Mom
    id: makeMoney ☐ Become a Millionaire
```

## **Advanced Shortcuts**
---


Use these flags

    - checklist -create a
    - checklist -read a
    - checklist -update a
    - checklist -delete a

they skip the intial question. [^1]

[^1]: Maybe some scerets ones to be added later.

## **Future Updates**

- [✔] Update the Readme
- [✔] Make Prettier
- [ ] Fix File Structure

Feel free to make suggestions.




