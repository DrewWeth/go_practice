/**
* Aaron
* Bobby    x
* Evan     x
* James
* Roy      x
* ......
* Tina     x
* ===========
* Evan, Bobby, Tina, Roy    >>>
*
* FriendList
*
* RecipientManager
*  toggle          O(1)
*  isSelected      O(1)
*  getRecipients   O(n)
*
*/

package main

import "fmt"

var p = fmt.Println

var hash map[string]*Recip
var end *Recip
var head *Recip

func main() {
	_ = Recip{"Roy", nil, nil}
	hash = make(map[string]*Recip)

	p(head)
	// toggle(r1.Name)
	// p(head)
	toggle("Drew")
	p(getRecipients())
	toggle("Abc")
	p(getRecipients())
	toggle("Abc")
	p(getRecipients())
	toggle("Drew")
	p(getRecipients())
	toggle("Drew1")
	toggle("Drew2")
	toggle("Drew3")
	p(getRecipients())
	toggle("Drew2")
	p(getRecipients())
}

type Recip struct {
	Name string
	prev *Recip
	next *Recip
}

func getRecipients() []string {
	temp := head
	// hash["Roy"] -> Linked List Object
	recips := []string{}
	if head == nil {
		return recips
	}

	for temp != nil {
		recips = append(recips, temp.Name)
		temp = temp.next
	}

	return recips
}

func toggle(new string) {
	recip := getRecip(new)

	if recip != nil { // In recip
		setRecip(new, nil)
		if recip.prev == nil { // Head in nil
			head = nil
			end = nil
			return
		}
		if recip.next == nil { // Main end pointer
			end = recip.prev
		}
		recip.prev.next = recip.next // Skip over

		} else { // Not in recip
			newRecip := &Recip{new, end, nil}
			setRecip(new, newRecip)

			if head == nil {
				head = newRecip
				end = head
				return
			}

			// end.prev.next = end
			end.next = newRecip
			end = end.next // Move end pointer to end
		}
	}

	func setRecip(name string, recip *Recip) {
		hash[name] = recip
	}

	func getRecip(name string) *Recip {
		obj, _ := hash[name]
		if obj != nil {
			return obj
		}
		return nil
	}

	func isSelected(name string) bool {
		if getRecip(name) != nil {
			return true
		}
		return false

	}
